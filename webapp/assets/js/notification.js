(function() {
	function registerWorker() {
		if ('serviceWorker' in navigator) {
			navigator.serviceWorker.register('/notification.js')
				.then(function(registration) {
					console.log('ServiceWorker registration successful with scope: ', registration.scope);
					initialise();
				})
				.catch(function(err) {
					console.log('ServiceWorker registration failed: ', err);
				});
		}
	}

	function initialise() {
		if (!('showNotification' in ServiceWorkerRegistration.prototype)) {  
			console.warn('Notifications aren\'t supported.');  
			return;  
		}

		if (Notification.permission === 'denied') {  
			console.warn('The user has blocked notifications.');
			return;  
		}

		if (!('PushManager' in window)) {  
			console.warn('Push messaging isn\'t supported.');  
			return;  
		}
		console.log("Is service worker ready?");
		navigator.serviceWorker.ready.then(function(serviceWorkerRegistration) {
			serviceWorkerRegistration.pushManager.subscribe({userVisibleOnly: true})
				.then(function(subscription) {
					if(!subscription) {
						return;
					}

					var endpoint = 'https://android.googleapis.com/gcm/send/';
					key = subscription.endpoint.replace(endpoint, '');
					console.log(endpoint, key);
				})
				.catch(function(err) {
					console.warn('Error during getSubscription()', err);  
				});
		});
	}

	document.getElementById('notify').addEventListener('click', function() {
		new Notification("This is an example notification", { icon: "/assets/images/icon-192x192.png", tag: "Notification" });
	});


	document.getElementById('push').addEventListener('click', function() {
		var req = new XMLHttpRequest();
		req.open('post', '/labs/notification/push', true);
		req.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
		req.send(JSON.stringify({registration_ids: [key]}));
	});


	registerWorker();
})();