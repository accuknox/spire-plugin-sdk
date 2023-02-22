// Code generated by protoc-gen-go-spire. DO NOT EDIT.

package agentstorev1

import (
	pluginsdk "github.com/accuknox/spire-plugin-sdk/pluginsdk"
	grpc "google.golang.org/grpc"
)

func AgentStoreServiceServer(server AgentStoreServer) pluginsdk.ServiceServer {
	return agentStoreServiceServer{AgentStoreServer: server}
}

type agentStoreServiceServer struct {
	AgentStoreServer
}

func (s agentStoreServiceServer) GRPCServiceName() string {
	return "spire.hostservice.server.agentstore.v1.AgentStore"
}

func (s agentStoreServiceServer) RegisterServer(server *grpc.Server) interface{} {
	RegisterAgentStoreServer(server, s.AgentStoreServer)
	return s.AgentStoreServer
}

type AgentStoreServiceClient struct {
	AgentStoreClient
}

func (c *AgentStoreServiceClient) IsInitialized() bool {
	return c.AgentStoreClient != nil
}

func (c *AgentStoreServiceClient) GRPCServiceName() string {
	return "spire.hostservice.server.agentstore.v1.AgentStore"
}

func (c *AgentStoreServiceClient) InitClient(conn grpc.ClientConnInterface) interface{} {
	c.AgentStoreClient = NewAgentStoreClient(conn)
	return c.AgentStoreClient
}
