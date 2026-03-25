<script lang="ts">
	import GuestLockSurface from "./guest-lock-surface.svelte";

	const bullets = [
		"Custom targets are shown as a gated capability, not hidden.",
		"Guest mode keeps allowlisted examples only.",
		"Arbitrary outbound URLs remain blocked until sign-in and validation.",
	];

	const highlights = [
		{ label: "Guest URLs", value: "Allowlisted" },
		{ label: "Custom domains", value: "Locked" },
		{ label: "Safety model", value: "Server enforced" },
	];
</script>

<GuestLockSurface
	eyebrow="Custom URLs"
	title="Custom outbound targets remain locked for guests"
	description="The shell still communicates where URL input lives, but guest sessions are limited to curated examples so the app cannot behave like a public proxy."
	badge="Outbound"
	summary="Authenticated users unlock server-validated destinations, redirect checks, body limits, and timeout controls. Guests stay on the allowlisted path."
	{bullets}
	{highlights}
	primaryActionLabel="Unlock custom targets"
	primaryActionHref="/app"
	secondaryActionLabel="Review API contract"
	secondaryActionHref="/docs"
	tone="rose"
/>
