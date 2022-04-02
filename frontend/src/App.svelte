<script>
	import ButtonComponent from "./ButtonComponent.svelte";

	export let name;
	let result = "stop";
	let active = false;
	let color = "red";


	async function status() {
		const res = await fetch("/status", {
			method: "GET",
		});

		const json = await res.text();

		result = json;
		active = result == "start";
		color = active ? "green" : "red";
	}
	status()
	
	async function startStop() {
		const path = result == "start" ? "/stop" : "/start";
		console.log(path);
		const res = await fetch(path, {
			method: "GET",
		});

		const json = await res.text();

		result = json;
		active = result == "start";
		color = active ? "green" : "red";
	}
</script>

<main>
	<h1>HOLA {name}!</h1>
	<h2>Calefaccion del Pueblo</h2>

	<p>Da al boton para parar o encender la caldera</p>

	<ButtonComponent on:click={startStop} symbol={result} bgColor={color} />
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}
	h2 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 2em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
