<script>
	import Nested from './Nested.svelte'
	let string = 'this string contains some <strong>HTML!!</strong>'

	let count = 0;
	$: doubled = count * 2;
	function incrementCount() {
		count++;
	}
</script>

<p>This is a paragraph.</p>
<p>{@html string}</p>
<Nested></Nested>
<button on:click={incrementCount}>
	Clicked {count} {count == 1 ? 'time': 'times'}
</button>
<p>{count} dubled id {doubled}</p>

<style>
	p {
		color: purple;
		font-family: 'Comic Sans MS', cursive;
		font-size: 2em;
	}
</style>