
document.addEventListener("DOMContentLoaded", () => {

	const header = document.querySelector('header');

	let ul = header.querySelector('ul');
	if (ul !== null) {

		for (let c = 0; c < ul.childNodes.length; c++) {

			let node = ul.childNodes[c];
			if (node.nodeType === 3 && node.nodeValue.trim() === '') {
				node.parentNode.removeChild(node);
				c--;
			}

		}

	}

});
