
document.addEventListener("DOMContentLoaded", () => {

	const STATISTICS = document.querySelector('main #statistics');
	const TABLE      = document.querySelector('main table tbody');

	const renderRepository = (owner, repository) => {

		let id   = owner + "/" + repository.name;
		let html = "";

		html += "<tr>";

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

	const renderStatistics = (users, organizations, repositories) => {

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

		return text.join(", ");

	};

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

		STATISTICS.innerHTML = renderStatistics(statistics.users, statistics.organizations, statistics.repositories);
		TABLE.innerHTML = table_rows.join('');

	});

});

