<script lang="ts">
	import GuestLockSurface from "./guest-lock-surface.svelte";

	const bullets = [
		"Environment variables remain visible as a product concept.",
		"Secret storage and scoped overrides are authenticated features.",
		"Guests can still inspect where values will be applied in the request flow.",
	];

	const highlights = [
		{ label: "Secret storage", value: "Locked" },
		{ label: "Variable editor", value: "Preview only" },
		{ label: "Scope", value: "Per account" },
	];
</script>

<GuestLockSurface
	eyebrow="Environment variables"
	title="Env vars stay locked for guest sessions"
	description="The surface remains part of the workspace so the product feels complete, but scoped secrets and reusable values are only available after authentication."
	badge="Variables"
	summary="Guests can inspect where request variables belong, while authenticated users get persistent, account-scoped values and safer request composition."
	{bullets}
	{highlights}
	primaryActionLabel="Sign in for variables"
	primaryActionHref="/app"
	secondaryActionLabel="Read the quick start"
	secondaryActionHref="/docs"
	tone="sand"
/>
