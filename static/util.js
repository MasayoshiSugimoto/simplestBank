function createClient(handler) {
	const request = new XMLHttpRequest()
	request.onreadystatechange = responseHandler(request, handler, console.log)
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

function responseHandler(request, resolve, reject) {
	return event => {
		if (request.readyState == XMLHttpRequest.DONE) {
			console.log(request.responseText)
			if (request.status === 200) {
				resolve(JSON.parse(request.responseText))
			} else {
				reject("createClient: request failure")
			}
		}
	}
}

function login() {
	const request = new XMLHttpRequest()
	request.onreadystatechange = responseHandler(
		request,
		console.log,
		console.log
	)
	request.open("POST", "/login", true)
	request.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")

	const options = new URLSearchParams()
	options.append("phone_number", document.getElementById("phone_number").value)

	request.send(options)
}
