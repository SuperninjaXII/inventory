let schedule_name = document.querySelector("#name");
let endTime = document.querySelector("#end");
let startTime = document.querySelector("#start");
getProducts = (name, start, end) => {
	fetch("/uploadschedule", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify([
			{
				name: name.value,
				start_time: `${start.value}`,
				end_time: `${end.value}`,
			},
		]),
	})
		.then((res) => {
			return res.json();
		})
		.then((data) => {
			console.log(JSON.stringify(data));
		});
	location.reload();
	alert("schedule added")
};
document
	.querySelector("#addSchedule")
	.addEventListener("click", () =>
		getProducts(schedule_name, endTime, startTime),
	);
