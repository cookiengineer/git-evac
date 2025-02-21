
document.addEventListener("DOMContentLoaded", () => {

	const ELEMENTS = {
		dialog: document.querySelector("dialog"),
		toggle: document.querySelector("main table thead input[type=\"checkbox\"]"),
		table:  document.querySelector("main table tbody"),
		footer: document.querySelector("footer")
	};

	const MESSAGES = {
		statistics: document.querySelector("main [data-message=\"statistics\"]"),
		selections: document.querySelector("footer [data-message=\"selections\"]")
	};

	const renderSelectionsMessage = (settings) => {

		MESSAGES.selections.innerHTML = "Selected " + Object.keys(settings).length + " Repositories";

	};

	const renderStatisticsMessage = (users, organizations, repositories) => {

		let text = [];

		if (users > 1 || users == 0) {
			text.push(users + " Users")
		} else {
			text.push(users + " User")
		}

		if (organizations > 1 || organizations == 0) {
			text.push(organizations + " Organizations")
		} else {
			text.push(organizations + " Organization")
		}

		if (repositories > 1 || repositories == 0) {
			text.push(repositories + " Repositories")
		} else {
			text.push(repositories + " Repository")
		}

		MESSAGES.statistics.innerHTML = text.join(", ");

	};

	if (ELEMENTS.toggle !== null) {

		ELEMENTS.toggle.onchange = () => {

			let is_checked = ELEMENTS.toggle.checked === true;

			Array.from(ELEMENTS.table.querySelectorAll("tr")).forEach((row) => {
				row.setAttribute("data-select", is_checked);
				row.querySelector("input[type=\"checkbox\"]").checked = is_checked;
				renderFooter();
			});

		};

	}

	fetch("/api/index").then((response) => {
		return response.json();
	}).then((data) => {

		let statistics = {
			users:         0,
			organizations: 0,
			repositories:  0
		};

		Object.keys(data.users).forEach((user_name) => {

			statistics.users++;

			Object.keys(data.users[user_name].repositories).forEach((repo_name) => {
				statistics.repositories++;
			});

		});

		Object.keys(data.organizations).forEach((orga_name) => {

			statistics.organizations++;

			Object.keys(data.organizations[orga_name].repositories).forEach((repo_name) => {
				statistics.repositories++;
			});

		});

		renderStatisticsMessage(statistics.users, statistics.organizations, statistics.repositories);

	});

});

