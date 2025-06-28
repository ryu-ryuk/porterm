// package main

// import (
// 	tea "github.com/charmbracelet/bubbletea"
// 	"porterm/model"
// 	"os"
// )

//	func main() {
//		p := tea.NewProgram(model.New())
//		if _, err := p.Run(); err != nil {
//			println("Error running program:", err)
//			os.Exit(1)
//		}
//	}
//
// -------------------------------------------
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"porterm/model"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
)

const (
	host        = "0.0.0.0"
	port        = 22
	hostKeyPath = ".ssh/term_host_rsa"
)

func main() {
	log.SetOutput(os.Stderr)

	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(hostKeyPath),
		wish.WithMiddleware(handleSSHConnection()),
	)
	if err != nil {
		log.Fatalf("could not start ssh server: %v", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("starting porterm ssh server on %s:%d", host, port)
	go func() {
		if err = s.ListenAndServe(); err != nil && err != ssh.ErrServerClosed {
			log.Fatalf("ssh server failed: %v", err)
		}
	}()

	<-done
	log.Println("stopping porterm ssh server gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("ssh server shutdown failed: %v", err)
	}
	log.Println("porterm ssh server stopped.")
}

func handleSSHConnection() wish.Middleware {
	return func(next ssh.Handler) ssh.Handler {
		return func(s ssh.Session) {
			// manually set TERM
			_ = os.Setenv("TERM", "xterm-256color")

			// apply all session env vars (important for TERM, LANG, etc)
			for _, env := range s.Environ() {
				parts := strings.SplitN(env, "=", 2)
				if len(parts) == 2 {
					_ = os.Setenv(parts[0], parts[1])
				}
			}

			p := tea.NewProgram(
				model.New(),
				tea.WithInput(s),
				tea.WithOutput(s),
				tea.WithAltScreen(),
				tea.WithMouseCellMotion(),
			)

			if _, err := p.Run(); err != nil {
				log.Printf("bubbletea error in session %s: %v", s.RemoteAddr(), err)
			}

			next(s)
		}
	}
}
