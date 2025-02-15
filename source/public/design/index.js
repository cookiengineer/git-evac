
document.addEventListener("DOMContentLoaded", () => {

	const header = document.querySelector('header');
	const footer = document.querySelector('footer');

	const cleanup = (node) => {

		for (let c = 0; c < node.childNodes.length; c++) {

			let child = node.childNodes[c];
			if (child.nodeType === 3 && child.nodeValue.trim() === '') {
				child.parentNode.removeChild(child);
				c--;
			}

		}

	};

	let ul = header.querySelector('ul');
	if (ul !== null) {
		cleanup(ul);
	}

	if (footer !== null) {
		cleanup(footer);
	}

});
