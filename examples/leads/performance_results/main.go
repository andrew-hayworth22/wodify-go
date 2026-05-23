// Command leads/performance_results demonstrates listing performance results for
// a lead using the Wodify Go SDK.
//
// Usage:
//
//	export WODIFY_API_KEY=your_api_key
//	make leads-performance_results
package main

import (
	"context"
	"fmt"
	"log"

	wodify "github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/leads"
	"github.com/joho/godotenv"
)

const leadID = 2191681
const componentID = 1234567

func main() {
	_ = godotenv.Load()
	ctx := context.Background()

	// Instantiate the Wodify client.
	wc, err := wodify.New()
	if err != nil {
		log.Fatalf("creating wodify client: %v\n", err)
	}

	// Fetch a lead's performance results.
	results, err := wc.Leads.ListPerformanceResults(ctx, leadID, leads.NewPerformanceResultListRequest(
		wodify.NewPaginationRequest(1, 10),
	))
	if err != nil {
		log.Fatalf("listing lead performance results: %v\n", err)
	}

	// Print the performance results.
	fmt.Printf("lead %d has %d performance result(s)\n", leadID, len(results.PerformanceResults))
	for _, result := range results.PerformanceResults {
		fmt.Printf("result: id=%s, class=%s, result type=%s, component ID=%d\n", result.ID, result.ClassName, result.ResultType, result.ComponentID)
	}

	// Fetch a lead's performance results by component ID.
	results, err = wc.Leads.ListPerformanceResultsByComponent(ctx, leadID, componentID, leads.NewPerformanceResultListRequest(
		wodify.NewPaginationRequest(1, 10),
	))
	if err != nil {
		log.Fatalf("listing performance results by component: %v\n", err)
	}

	// Print the lead's performance results by component ID.
	fmt.Printf("lead %d has %d performance result(s) with component ID =%d\n", leadID, len(results.PerformanceResults), componentID)
	for _, result := range results.PerformanceResults {
		fmt.Printf("result: id=%s, class=%s, result type=%s, component ID=%d\n", result.ID, result.ClassName, result.ResultType, result.ComponentID)
	}
}
