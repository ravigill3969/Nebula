package orchestrator

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gopkg.in/yaml.v3"

	pb "nebula/proto"
)

type WorkflowServer struct {
	pb.UnimplementedWorkflowServiceServer
}

type Workflow struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name      string `yaml:"name"`
	Service   string `yaml:"service"`
	Action    string `yaml:"action"`
	DependsOn string `yaml:"depends_on,omitempty"`
}

var workerRegistry = map[string]string{
	"user-service":    "localhost:6001",
	"billing-service": "localhost:6002",
	"email-service":   "localhost:6003",
}

func (s *WorkflowServer) SubmitWorkflow(ctx context.Context, req *pb.WorkflowDefinition) (*pb.WorkflowResponse, error) {
	fmt.Printf("📩 Received workflow: %s\n", req.Name)

	var wf Workflow
	if err := yaml.Unmarshal(req.YamlFile, &wf); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	id := uuid.New().String()

	fmt.Printf("✅ Workflow [%s] completed.\n", id)
	return &pb.WorkflowResponse{
		WorkflowId: id,
		Status:     "completed",
	}, nil
}

func ExecuteWorkflow(wf Workflow) {
	fmt.Printf("🚀 Starting workflow: %s\n", wf.Name)
	completed := make(map[string]bool)

	for len(completed) < len(wf.Steps){

	}
}


