<script lang="ts">
	import GuestLockSurface from "./guest-lock-surface.svelte";

	const bullets = [
		"History remains visible in the shell for guest orientation.",
		"Runs are ephemeral until an account is attached.",
		"Re-run and persist controls stay clearly disabled.",
	];

	const highlights = [
		{ label: "Guest retention", value: "Ephemeral" },
		{ label: "Signed-in retention", value: "Stored runs" },
		{ label: "Visibility", value: "Always shown" },
	];
</script>

<GuestLockSurface
	eyebrow="Request history"
	title="History persistence is locked until authentication"
	description="Visitors can see where recent runs will appear, but nothing is written to long-term storage until the session is tied to an account."
	badge="History rail"
	summary="This keeps the workspace honest: guests still understand the request timeline, while authenticated users get durable history and replay support."
	{bullets}
	{highlights}
	primaryActionLabel="Sign in to view history"
	primaryActionHref="/app"
	secondaryActionLabel="See workspace layout"
	secondaryActionHref="/docs"
	tone="slate"
/>
