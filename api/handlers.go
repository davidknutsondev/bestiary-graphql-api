package api

import (
	"net/http"

	"github.com/davidknutsondev/bestiary-graphql-api/pkg/graphql/schema" // Update with the correct import path
	"github.com/graphql-go/handler"
)

var schemaHandler = handler.New(&handler.Config{
	Schema:   &schema.BeastSchema,
	Pretty:   true,
	GraphiQL: false,
})

func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	schemaHandler.ServeHTTP(w, r)
}

func SandboxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(sandboxHTML)
}

var sandboxHTML = []byte(`
<!DOCTYPE html>
<html lang="en">
<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
<div id="sandbox" style="height:100vh; width:100vw;"></div>
<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
<script>
 new window.EmbeddedSandbox({
   target: "#sandbox",
   // Pass through your server href if you are embedding on an endpoint.
   // Otherwise, you can pass whatever endpoint you want Sandbox to start up with here.
   initialEndpoint: "http://localhost:8080/graphql",
 });
 // advanced options: https://www.apollographql.com/docs/studio/explorer/sandbox#embedding-sandbox
</script>
</body>

</html>`)
