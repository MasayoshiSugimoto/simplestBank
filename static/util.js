function createClient(handler) {
	const request = new XMLHttpRequest()
	request.onreadystatechange = event => {
		if (request.readyState == XMLHttpRequest.DONE) {
			if (request.status === 200) {
				handler(JSON.parse(request.responseText))
			} else {
				console.log("createClient: request failure")
			}
		}
	}
	request.open("POST", "/client/new", true)
	request.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")

	const options = new URLSearchParams()
	options.append("first_name", document.getElementById("first_name").value)
	options.append("last_name", document.getElementById("last_name").value)
	options.append("phone_number", document.getElementById("phone_number").value)

	request.send(options)
}

function onClientCreatePerson() {
	createClient(response =>
		document.getElementById("alert_component").innerText = response.Message
	)
}
