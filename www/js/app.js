function getAll(entity) {
	fetch('https://airlines-ffaas.netlify.app//api/' + entity)
	  .then((response) => response.json())
		.then((data) => {
			fetch('/template/list/' + entity + '.html')
				.then((response) => response.text())
				.then((template) => {
					var rendered = Mustache.render(template, data);
					console.log('template content');
					console.log(rendered);
					console.log(data);
					document.getElementById('content').innerHTML = rendered;
				});
		})
}

function getById(query, entity) {
	var params = new URLSearchParams(query);
	fetch('https://airlines-ffaas.netlify.app//api/' + entity + '/?id=' + params.get('id'))
	  .then((response) => response.json())
		.then((data) => {
			fetch('/template/detail/' + entity + '.html')
				.then((response) => response.text())
				.then((template) => {
					var rendered = Mustache.render(template, data);
					console.log('template content');
					console.log(rendered);
					console.log(data);
					document.getElementById('content').innerHTML = rendered;
				});
		})
}

function home() {
	fetch('/template/home.html')
		.then((response) => response.text())
		.then((template) => {
			var rendered = Mustache.render(template, {});
			document.getElementById('content').innerHTML = rendered;
		});
}

function init() {
	router = new Navigo(null, false, '#!');
	router.on({
		'/flights': function() {
			getAll('flights');
		},
		'/airports': function() {
			getAll('airports');
		},
		'/travelers': function() {
			getAll('travelers');
		},
		'/flightById': function(_, query) {
			getById(query, 'flights');
		},
		'/airportById': function(_, query) {
			getById(query, 'airports');
		},
		'/travelerById': function(_, query) {
			getById(query, 'travelers');
		}
	});
	router.on(() => home());
	router.resolve();
}
