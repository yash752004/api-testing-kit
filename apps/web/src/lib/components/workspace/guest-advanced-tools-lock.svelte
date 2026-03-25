<script lang="ts">
	import GuestLockSurface from "./guest-lock-surface.svelte";

	const bullets = [
		"Advanced utilities remain in place for product completeness.",
		"Code snippets, export tools, and account-level extras are gated.",
		"Guests can discover the feature set without triggering the workflows.",
	];

	const highlights = [
		{ label: "Snippets", value: "Preview only" },
		{ label: "Advanced tools", value: "Locked" },
		{ label: "Unlock path", value: "Authenticated" },
	];
</script>

<GuestLockSurface
	eyebrow="Advanced tools"
	title="Advanced workspace tools are intentionally gated"
	description="The lower-priority surfaces still appear in the product so the shell feels real, but the powerful utilities stay behind authentication until the safety model is in place."
	badge="Utility drawer"
	summary="This covers tools like snippets, history depth, environment variables, and other power-user surfaces without making the guest experience feel fake."
	{bullets}
	{highlights}
	primaryActionLabel="Sign in for tools"
	primaryActionHref="/app"
	secondaryActionLabel="Open the docs"
	secondaryActionHref="/docs"
	tone="emerald"
/>
