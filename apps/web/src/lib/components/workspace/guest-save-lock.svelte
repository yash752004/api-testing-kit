<script lang="ts">
	import GuestLockSurface from "./guest-lock-surface.svelte";

	const bullets = [
		"Collection save and duplicate actions remain visible.",
		"Permanent request storage unlocks after sign-in.",
		"Guest sessions can review the save workflow without writing data.",
	];

	const highlights = [
		{ label: "Guest save", value: "Preview only" },
		{ label: "Authenticated mode", value: "Persistent" },
		{ label: "Surface type", value: "Card + overlay" },
	];
</script>

<GuestLockSurface
	eyebrow="Save requests"
	title="Saving requests is intentionally locked for guests"
	description="The workspace still shows the save affordance, but permanent storage is reserved for authenticated users so the product never feels like an open scratchpad."
	badge="Save control"
	summary="Guests can understand the save flow and its placement in the shell, while the backend keeps storage and collection management behind login."
	{bullets}
	{highlights}
	primaryActionLabel="Sign in to save"
	primaryActionHref="/app"
	secondaryActionLabel="Open collections docs"
	secondaryActionHref="/docs"
	tone="emerald"
/>
