package main

import (
	"context" // new import
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"porterm/model"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
)

const (
	host = "0.0.0.0" // listen on all interfaces
	port = 2022
	// path to your ssh host key. generate this file:\
	// doc:
	// ssh-keygen -t rsa -b 4096 -f .ssh/term_host_rsa -N ""
	hostKeyPath = ".ssh/term_host_rsa"
)

func main() {
	// direct log output to stderr, useful for systemd
	log.SetOutput(os.Stderr)

	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(hostKeyPath),
		wish.WithMiddleware(
			handleSSHConnection(),
		),
	)
	if err != nil {
		log.Fatalf("could not start ssh server: %v", err)
	}

	// shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("starting porterm ssh server on %s:%d", host, port)
	go func() {
		if err = s.ListenAndServe(); err != nil {
			// this error is expected on shutdown, so check if it's not a server closed error
			if err != ssh.ErrServerClosed {
				log.Fatalf("ssh server failed: %v", err)
			}
		}
	}()

	<-done
	log.Println("stopping porterm ssh server gracefully...")
	// create a context with a timeout for proper shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("ssh server shutdown failed: %v", err)
	}
	log.Println("porterm ssh server stopped.")
}

// returns a wish.Middleware that runs the bubble tea program.
func handleSSHConnection() wish.Middleware {
	return func(next ssh.Handler) ssh.Handler {
		return func(s ssh.Session) {
			// create a new bubble tea program for each ssh session
			p := tea.NewProgram(
				model.New(), // calls model.New()
				tea.WithInput(s),
				tea.WithOutput(s),
			)

			// run the bubble tea program. the program will exit when 'q' or 'ctrl+c' is pressed.
			_, err := p.Run()
			if err != nil {
				log.Printf("bubble tea program exited with error for session %s: %v", s.RemoteAddr(), err)
			}

			next(s)
		}
	}
}
