package client

import (
	"bufio"
	"context"
	"fmt"
	challenge "github.com/gorobot-nz/go-grpc-task/pkg/gen/pkg/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	TIMER      = "timer"
	SHORT_LINK = "short"
	META       = "meta"
	QUIT       = "quit"
)

type Client struct {
	challengeServiceClient challenge.ChallengeServiceClient
}

func NewClient(port string) *Client {
	conn, err := grpc.Dial(port, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to PetStoreService on %s: %w", port, err)
	}
	challengeClient := challenge.NewChallengeServiceClient(conn)
	return &Client{challengeServiceClient: challengeClient}
}

func (c *Client) Run() error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var line string
		fmt.Print("Enter command > ")
		if scanner.Scan() {
			line = scanner.Text()
		}
		args := strings.Split(line, " ")
		switch args[0] {
		case TIMER:
			name := args[1]
			time, _ := strconv.Atoi(args[2])
			frequency, _ := strconv.Atoi(args[3])
			err := c.startTimer(name, time, frequency)
			if err != nil {
				return err
			}
		case META:
			metadata := args[1]
			err := c.readMetadata(metadata)
			if err != nil {
				return err
			}
		case SHORT_LINK:
			link := args[1]
			err := c.shortLink(link)
			if err != nil {
				return err
			}
		case QUIT:
			return nil
		default:
			c.help()
		}
	}
}

func (c *Client) shortLink(link string) error {
	resp, err := c.challengeServiceClient.MakeShortLink(context.Background(), &challenge.Link{Data: link})
	if err != nil {
		return err
	}
	log.Println(resp.GetData())
	return nil
}

func (c *Client) readMetadata(metadata string) error {
	resp, err := c.challengeServiceClient.ReadMetadata(context.Background(), &challenge.Placeholder{Data: metadata})
	if err != nil {
		return err
	}
	log.Println(resp.GetData())
	return nil
}

func (c *Client) startTimer(name string, time int, frequency int) error {
	stream, err := c.challengeServiceClient.StartTimer(context.Background(), &challenge.Timer{Name: name, Seconds: int64(time), Frequency: int64(frequency)})
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", c.challengeServiceClient, err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}

		fmt.Println(res.GetSeconds())
	}
	return nil
}

func (c *Client) help() {
	fmt.Println("Example usage")
	fmt.Println("\tmeta data(string)")
	fmt.Println("\tshort link(string)")
	fmt.Println("\ttimer name(string) time(int) frequency(int)")
	fmt.Println("\tquit")
	fmt.Println("Good luck")
}
