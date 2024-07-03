package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/rubaiat-hossain/terraform-provider-platformsh/internal/provider"
)

func main() {
	err := providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "local.provider/rhs/platformsh",
	})
	if err != nil {
		log.Fatal(err)
	}
}
