<script lang="ts">
	import AppShellShell from "$lib/components/app-shell/app-shell-shell.svelte";

	let { children } = $props();
</script>

<AppShellShell>
	{@render children()}
</AppShellShell>
