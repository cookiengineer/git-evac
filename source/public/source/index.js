
document.addEventListener("DOMContentLoaded", () => {

	const ELEMENTS = {
		toggle: document.querySelector("main table thead input[type=\"checkbox\"]"),
		table:  document.querySelector("main table tbody"),
		footer: document.querySelector("footer")
	};

	const MESSAGES = {
		statistics: document.querySelector("main [data-message=\"statistics\"]"),
		selections: document.querySelector("footer [data-message=\"selections\"]")
	};

	const openDialog = (settings) => {

		// settings is map[repository] = action
		console.warn("TODO: Open Dialog");

	};

	const renderFooter = () => {

		let selected = 0;
		let settings = {};

		Array.from(ELEMENTS.table.querySelectorAll("tr[data-select=\"true\"]")).forEach((row) => {

			let id = row.getAttribute("data-id");
			if (id !== "") {

				let button = row.querySelector("button[data-action]");
				if (button !== null) {

					settings[id] = button.getAttribute("data-action");
					selected++;

				}

			}


		});

		renderFooterActions(settings);
		renderSelectionsMessage(settings);

	};

	const renderFooterActions = (settings) => {

		// settings is map[repository] = action

		let needs_fix    = [];
		let needs_commit = [];
		let can_pull_and_push = [];

		Object.keys(settings).forEach((identifier) => {

			let action = settings[identifier];

			if (action === "fix") {
				needs_fix.push(identifier);
			} else if (action === "commit") {
				needs_commit.push(identifier);
			} else if (action === "push" || action === "pull") {
				can_pull_and_push.push(identifier);
			}

		});

		let buttons = [];

		if (needs_fix.length > 0) {
			buttons.push("<button data-action=\"fix\">Fix " + needs_fix.length + "</button>");
			buttons.push("<button data-action=\"commit\" disabled>Commit " + needs_commit.length + "</button>");
			buttons.push("<button data-action=\"pull\" disabled>Pull " + can_pull_and_push.length + "</button>");
			buttons.push("<button data-action=\"push\" disabled>Push " + can_pull_and_push.length + "</button>");
		} else if (needs_commit.length > 0) {
			buttons.push("<button data-action=\"commit\">Commit " + needs_commit.length + "</button>");
			buttons.push("<button data-action=\"pull\" disabled>Pull " + can_pull_and_push.length + "</button>");
			buttons.push("<button data-action=\"push\" disabled>Push " + can_pull_and_push.length + "</button>");
		} else if (can_pull_and_push.length > 0) {
			buttons.push("<button data-action=\"pull\">Pull " + can_pull_and_push.length + "</button>");
			buttons.push("<button data-action=\"push\">Push " + can_pull_and_push.length + "</button>");
		}

		let actions_element = ELEMENTS.footer.querySelector("div:last-of-type");
		if (actions_element !== null) {

			if (buttons.length > 0) {
				actions_element.innerHTML = buttons.join("");
			} else {
				actions_element.innerHTML = "<button data-action=\"refresh\">Refresh</button>";
			}
		}

	};

	const renderRepository = (owner, repository) => {

		let id   = owner + "/" + repository.name;
		let html = "";

		html += "<tr data-id=\"" + id + "\" data-select=\"false\">";

		html += "<td><input type=\"checkbox\" data-id=\"" + id + "\" name=\"" + id + "\"/></td>";
		html += "<td>" + owner + "/" + repository.name + "</td>";

		html += "<td>" + Object.keys(repository.remotes).sort().map((remote) => {

			if (repository.current_remote == remote) {
				return "<em>" + remote + "</em>";
			} else {
				return "<span>" + remote + "</span>";
			}

		}).join(" ") + "</td>";

		html += "<td>" + Array.from(repository.branches).sort().map((branch) => {

			if (repository.current_branch == branch) {
				return "<em>" + branch + "</em>";
			} else {
				return "<span>" + branch + "</span>";
			}

		}).join(" ") + "</td>";

		html += "<td>";
		if (repository.has_remote_changes == true) {
			html += "<em>remote changes</em>";
		} else if (repository.has_local_changes == true) {
			html += "<em>local changes</em>";
		} else {
			html += "";
		}
		html += "</td>";

		html += "<td>"
		if (repository.has_remote_changes == true) {
			html += "<button data-action=\"fix\">Fix</button>";
		} else if (repository.has_local_changes == true) {
			html += "<button data-action=\"commit\">Commit</button>";
		} else {
			html += "<button data-action=\"pull\">Pull</button>";
			html += "<button data-action=\"push\">Push</button>";
		}
		html += "</td>";

		html += "</tr>";

		return html;

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

	if (ELEMENTS.table !== null) {

		ELEMENTS.table.onclick = (event) => {

			let element = event.target;
			let tagname = element.tagName.toLowerCase();

			if (tagname == "input") {

				let row = element.parentNode.parentNode;
				let is_checked = element.checked;
				if (is_checked === true) {
					row.setAttribute("data-select", true);
				} else {
					row.setAttribute("data-select", false);
				}

				renderFooter();

			} else if (tagname == "button") {

				let row      = element.parentNode.parentNode;
				let id       = row.getAttribute("data-id");
				let action   = element.getAttribute("data-action");
				let settings = {
					[id]: action
				};

				openDialog(settings);

			} else if (tagname == "td") {

				let row = element.parentNode;
				let is_checked = row.getAttribute("data-select");
				if (is_checked === "true") {
					row.setAttribute("data-select", false);
					row.querySelector("input[type=\"checkbox\"]").checked = false;
				} else {
					row.setAttribute("data-select", true);
					row.querySelector("input[type=\"checkbox\"]").checked = true;
				}

				renderFooter();

			} else if (tagname == "em" || tagname == "span") {

				let row = element.parentNode.parentNode;
				let is_checked = row.getAttribute("data-select");
				if (is_checked === "true") {
					row.setAttribute("data-select", false);
					row.querySelector("input[type=\"checkbox\"]").checked = false;
				} else {
					row.setAttribute("data-select", true);
					row.querySelector("input[type=\"checkbox\"]").checked = true;
				}

				renderFooter();

			}

		};

	}

	fetch("/api/index").then((response) => {
		return response.json();
	}).then((data) => {

		let table_rows = [];
		let statistics = {
			users:         0,
			organizations: 0,
			repositories:  0
		};

		Object.keys(data.users).forEach((user_name) => {

			statistics.users++;

			Object.keys(data.users[user_name].repositories).forEach((repo_name) => {

				let repository = data.users[user_name].repositories[repo_name];
				table_rows.push(renderRepository("@" + user_name, repository));
				statistics.repositories++;

			});

		});

		Object.keys(data.organizations).forEach((orga_name) => {

			statistics.organizations++;

			Object.keys(data.organizations[orga_name].repositories).forEach((repo_name) => {

				let repository = data.organizations[orga_name].repositories[repo_name];
				table_rows.push(renderRepository(orga_name, repository));
				statistics.repositories++;

			});

		});

		renderStatisticsMessage(statistics.users, statistics.organizations, statistics.repositories);
		ELEMENTS.table.innerHTML = table_rows.join("");

	});

});

