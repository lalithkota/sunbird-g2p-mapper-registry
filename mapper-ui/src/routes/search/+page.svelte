<script>
	let value = "";
	let type = "id";
	function onSearch() {
		fetch("https://g2p-financial-wrapper.xiv.in/mapper-service/v0.1.0/mapper/search", {
			method: 'post',
			headers: {
				'content-type': 'application/json'
			},
			body: JSON.stringify({
				"signature": "Signature:  namespace=\"g2p\", kidId=\"{sender_id}|{unique_key_id}|{algorithm}\", algorithm=\"ed25519\", created=\"1606970629\", expires=\"1607030629\", headers=\"(created) (expires) digest\", signature=\"Base64(signing content)",
				"header": {
					"version": "0.1.0",
					"message_id": "123456789020211216223812",
					"message_ts": "2022-12-04T18:01:07+00:00",
					"action": "search",
					"sender_id": "registry.example.org",
					"sender_uri": "https://registry.sender.org/g2p/callback/on-disburse",
					"receiver_id": "pymts.example.org",
					"total_count": 21800,
					"is_encrypted": true
				},
				"message": {
					"transaction_id": "12345678901234567000",
					"search_request": {
						"reference_id": "12345678901234567890",
						"request_type": "link",
						"attribute_type": type,
						"attribute_value": value,
						"locale": "eng"
					}
				}
			})
		})
				.then(response => response.json())
				.then(data => {
					console.log(data);
				}).catch(error => {
			console.log(error);
			return [];
		});
	}
</script>
<div class="text-column">
	<div>
		<span>Search type:</span>
		<select bind:value={type}>
			<option>id</option>
			<option>fa</option>
		</select>
	</div>
	<div>
		<span>ID:</span>
		<input bind:value={value}/>
	</div>

	<button on:click={onSearch}>Search</button>
</div>
<style>
	div {
		margin-top: 1rem;
	}
</style>
