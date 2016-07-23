function doLogin() {
 	console.log("Inside doLogin()");
 	var username = $('#usr').val();
 	var password = $('#pwd').val();
 	var params = {};
 	params.loginName = username;
 	params.userPassword = password;
 	var jsonRequest = JSON.stringify(params);
 	var customerId;
 	$.ajax({
		url: "/UserLogin/Login",
		headers: { 'UserTimeZone': getTimeZone() },
		type: "POST",
		async: true,
		dataType: "json",
		data: jsonRequest,
		success: function(response) {
			console.log("Inside ajax success");
			var stringResponse = JSON.stringify(response);
			var parsedResponse = JSON.parse(stringResponse);
			customerId = parsedResponse.customerId;
			var token = parsedResponse.securityToken;
			console.log(customerId);
			console.log(token);
			localStorage.setItem("customerId", customerId);
			localStorage.setItem("token", token);
			console.log(localStorage.getItem("customerId"));
			console.log(localStorage.getItem("token"));
			location.href = "/Presentation/" + customerId;
		},
		failure: function(errorMsg) {
			console.log("AJAX failure");
		},
		error: function(jqXHR, status, error) {
			console.log("AJAX error");
			console.log(jqXHR.status);
			$('#responseLabel').html("This user does not exist. Please try again.");
		}
	});
}
		
function getTimeZone() {
    var offset = new Date().getTimezoneOffset(), o = Math.abs(offset);
	return "UTC" + (offset < 0 ? "+" : "-") + ("00" + Math.floor(o / 60)).slice(-2) + ":" + ("00" + (o % 60)).slice(-2);
}