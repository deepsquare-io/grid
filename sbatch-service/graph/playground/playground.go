// Copyright (C) 2024 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package playground

import (
	"html/template"
	"net/http"
	"net/url"
)

var apolloSandboxPage = template.Must(template.New("ApolloSandbox").Parse(`<!doctype html>
<html>

<head>
  <meta charset="utf-8">
  <title>{{.title}}</title>
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <link rel="icon" href="https://embeddable-sandbox.cdn.apollographql.com/_latest/public/assets/favicon-dark.png">
	<style>
	body {
		margin: 0;
		overflow: hidden;
	}
</style>
</head>

<body>
  <div style="width: 100vw; height: 100vh;" id='embedded-sandbox'></div>
  <!-- NOTE: New version available at https://embeddable-sandbox.cdn.apollographql.com/ -->
  <script rel="preload" as="script" crossorigin="anonymous" integrity="{{.mainSRI}}" type="text/javascript" src="https://embeddable-sandbox.cdn.apollographql.com/7212121cad97028b007e974956dc951ce89d683c/embeddable-sandbox.umd.production.min.js"></script>
  <script>
{{- if .endpointIsAbsolute}}
	const url = {{.endpoint}};
{{- else}}
	const url = location.protocol + '//' + location.host + {{.endpoint}};
{{- end}}
	<!-- See https://www.apollographql.com/docs/graphos/explorer/sandbox/#options -->
  new window.EmbeddedSandbox({
    target: '#embedded-sandbox',
    initialEndpoint: url,
		persistExplorerState: true,
		initialState: {
			includeCookies: true,
			pollForSchemaUpdates: false,
		}
  });
  </script>
</body>

</html>`))

// endpointHasScheme checks if the endpoint has a scheme.
func endpointHasScheme(endpoint string) bool {
	u, err := url.Parse(endpoint)
	return err == nil && u.Scheme != ""
}

// ApolloSandboxHandler responsible for setting up the altair playground
func ApolloSandboxHandler(title, endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := apolloSandboxPage.Execute(w, map[string]interface{}{
			"title":              title,
			"endpoint":           endpoint,
			"endpointIsAbsolute": endpointHasScheme(endpoint),
			"mainSRI":            "sha256-ldbSJ7EovavF815TfCN50qKB9AMvzskb9xiG71bmg2I=",
		})
		if err != nil {
			panic(err)
		}
	}
}
