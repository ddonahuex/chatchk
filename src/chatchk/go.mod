module nethopper.io/chatchk

go 1.24

require (
	nethopper.io/admin v0.0.0
	nethopper.io/ingest v0.0.0
	nethopper.io/knowledge v0.0.0
	nethopper.io/ollama v0.0.0
	nethopper.io/open_webui v0.0.0
	nethopper.io/prompts v0.0.0
)

require nethopper.io/utils v0.0.0 // indirect

replace nethopper.io/ingest => ../ingest

replace nethopper.io/knowledge => ../knowledge

replace nethopper.io/prompts => ../prompts

replace nethopper.io/admin => ../admin

replace nethopper.io/open_webui => ../open_webui

replace nethopper.io/ollama => ../ollama

replace nethopper.io/utils => ../utils
