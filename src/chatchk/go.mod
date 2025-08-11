module ddonahuex.io/chatchk

go 1.24

require (
	ddonahuex.io/admin v0.0.0
	ddonahuex.io/ingest v0.0.0
	ddonahuex.io/knowledge v0.0.0
	ddonahuex.io/ollama v0.0.0
	ddonahuex.io/open_webui v0.0.0
	ddonahuex.io/prompts v0.0.0
)

require ddonahuex.io/utils v0.0.0 // indirect

replace ddonahuex.io/ingest => ../ingest

replace ddonahuex.io/knowledge => ../knowledge

replace ddonahuex.io/prompts => ../prompts

replace ddonahuex.io/admin => ../admin

replace ddonahuex.io/open_webui => ../open_webui

replace ddonahuex.io/ollama => ../ollama

replace ddonahuex.io/utils => ../utils
