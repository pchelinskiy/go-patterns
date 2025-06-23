package adapter

import "fmt"

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(comp Computer) {
	fmt.Println("client inserts Lightning connector into computer")
	comp.InsertIntoLightningPort()
}
