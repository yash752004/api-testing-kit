<script lang="ts">
	import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "$lib/components/ui/card/index.js";
	import Badge from "$lib/components/ui/badge/badge.svelte";
	import { Button } from "$lib/components/ui/button/index.js";
	import {
		Tabs,
		TabsContent,
		TabsList,
		TabsTrigger,
	} from "$lib/components/ui/tabs/index.js";
	import Separator from "$lib/components/ui/separator/separator.svelte";
	import PlayIcon from "@lucide/svelte/icons/play";
	import LockIcon from "@lucide/svelte/icons/lock";
	import FileJsonIcon from "@lucide/svelte/icons/file-json";
	import InboxIcon from "@lucide/svelte/icons/inbox";
	import Clock3Icon from "@lucide/svelte/icons/clock-3";
	import ShieldIcon from "@lucide/svelte/icons/shield";
	import LayoutGridIcon from "@lucide/svelte/icons/layout-grid";
	import ArrowRightIcon from "@lucide/svelte/icons/arrow-right";

	const requestFields = [
		{ label: "Method", value: "GET" },
		{ label: "URL", value: "https://jsonplaceholder.typicode.com/posts/1" },
		{ label: "Auth", value: "None" },
	];

	const responseHeaders = [
		["content-type", "application/json; charset=utf-8"],
		["cache-control", "max-age=43200"],
		["x-powered-by", "Express"],
	];

	const prettyResponse = `{
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}`;

	const rawResponse = `{"id":1,"title":"delectus aut autem","completed":false}`;
</script>

<div class="grid gap-4 xl:grid-cols-[1.25fr_0.95fr]">
	<Card class="border-border/80 shadow-sm">
		<CardHeader class="gap-3">
			<div class="flex items-center justify-between gap-3">
				<div>
					<CardTitle>Request builder</CardTitle>
					<CardDescription>Static shell only. No request execution is wired yet.</CardDescription>
				</div>
				<Badge>Guest-safe</Badge>
			</div>

			<div class="flex flex-wrap items-center gap-2">
				{#each requestFields as field}
					<div class="rounded-full border border-border/70 bg-panel-soft px-3 py-1.5 text-xs">
						<span class="text-muted-foreground">{field.label}:</span>
						<span class="ml-1 font-medium text-foreground">{field.value}</span>
					</div>
				{/each}
			</div>
		</CardHeader>

		<CardContent class="space-y-4">
			<div class="rounded-[22px] border border-border/70 bg-shell p-4">
				<div class="flex flex-wrap items-center justify-between gap-3">
					<div class="flex items-center gap-2">
						<Badge variant="secondary">GET</Badge>
						<p class="font-mono text-sm text-foreground">/posts/1</p>
					</div>
					<Button variant="default" size="sm">
						<PlayIcon class="size-4" />
						Send
					</Button>
				</div>

				<div class="mt-4 rounded-2xl border border-border/70 bg-white p-4">
					<Tabs value="params" class="gap-4">
						<TabsList>
							<TabsTrigger value="params">Params</TabsTrigger>
							<TabsTrigger value="headers">Headers</TabsTrigger>
							<TabsTrigger value="body">Body</TabsTrigger>
							<TabsTrigger value="auth">Auth</TabsTrigger>
						</TabsList>

						<TabsContent value="params" class="space-y-2">
							<div class="grid grid-cols-[1.1fr_1fr_0.6fr] gap-2 text-xs text-muted-foreground">
								<span>Key</span>
								<span>Value</span>
								<span>State</span>
							</div>
							<div class="grid grid-cols-[1.1fr_1fr_0.6fr] gap-2 rounded-2xl border border-border/70 bg-panel-soft px-3 py-2 text-sm">
								<span class="font-mono">userId</span>
								<span>1</span>
								<span>Enabled</span>
							</div>
							<div class="grid grid-cols-[1.1fr_1fr_0.6fr] gap-2 rounded-2xl border border-border/70 bg-panel-soft px-3 py-2 text-sm">
								<span class="font-mono">sort</span>
								<span>desc</span>
								<span>Optional</span>
							</div>
						</TabsContent>

						<TabsContent value="headers" class="space-y-2">
							<div class="rounded-2xl border border-border/70 bg-panel-soft px-3 py-2 text-sm">
								<div class="flex items-center justify-between gap-3">
									<span class="font-mono">accept</span>
									<span>application/json</span>
								</div>
							</div>
							<div class="rounded-2xl border border-border/70 bg-panel-soft px-3 py-2 text-sm">
								<div class="flex items-center justify-between gap-3">
									<span class="font-mono">x-demo-mode</span>
									<span>guest</span>
								</div>
							</div>
						</TabsContent>

						<TabsContent value="body" class="space-y-3">
							<div class="rounded-2xl border border-dashed border-border/80 bg-white p-4">
								<p class="text-sm font-medium text-foreground">Body editor placeholder</p>
								<p class="mt-1 text-xs leading-5 text-muted-foreground">
									JSON, raw, and form modes will live here once request editing is wired.
								</p>
							</div>
						</TabsContent>

						<TabsContent value="auth" class="space-y-3">
							<div class="flex items-center gap-3 rounded-2xl border border-border/70 bg-panel-soft p-4">
								<LockIcon class="size-4 text-muted-foreground" />
								<div>
									<p class="text-sm font-medium text-foreground">Guest mode locked</p>
									<p class="text-xs text-muted-foreground">
										Bearer tokens and custom auth flows unlock after sign-in.
									</p>
								</div>
							</div>
						</TabsContent>
					</Tabs>
				</div>
			</div>
		</CardContent>
	</Card>

	<Card class="border-border/80 shadow-sm">
		<CardHeader class="gap-3">
			<div class="flex items-center justify-between gap-3">
				<div>
					<CardTitle>Response viewer</CardTitle>
					<CardDescription>Preset response states for the workspace shell.</CardDescription>
				</div>
				<Badge variant="secondary">200 OK</Badge>
			</div>
			<div class="flex flex-wrap gap-2 text-xs text-muted-foreground">
				<span class="rounded-full border border-border/70 bg-panel-soft px-3 py-1">186 ms</span>
				<span class="rounded-full border border-border/70 bg-panel-soft px-3 py-1">1.2 KB</span>
				<span class="rounded-full border border-border/70 bg-panel-soft px-3 py-1">application/json</span>
			</div>
		</CardHeader>

		<CardContent class="space-y-4">
			<div class="rounded-[22px] border border-border/70 bg-shell p-4">
				<Tabs value="pretty" class="gap-4">
					<TabsList>
						<TabsTrigger value="pretty">Pretty</TabsTrigger>
						<TabsTrigger value="raw">Raw</TabsTrigger>
						<TabsTrigger value="headers">Headers</TabsTrigger>
					</TabsList>

					<TabsContent value="pretty" class="space-y-3">
						<pre class="overflow-x-auto rounded-2xl border border-border/70 bg-slate-950 px-4 py-4 text-xs leading-6 text-slate-100"><code>{prettyResponse}</code></pre>
					</TabsContent>

					<TabsContent value="raw" class="space-y-3">
						<div class="rounded-2xl border border-border/70 bg-panel-soft p-4 font-mono text-xs leading-6 text-foreground">
							{rawResponse}
						</div>
					</TabsContent>

					<TabsContent value="headers" class="space-y-2">
						{#each responseHeaders as [key, value]}
							<div class="flex items-center justify-between gap-4 rounded-2xl border border-border/70 bg-panel-soft px-3 py-2 text-sm">
								<span class="font-mono text-xs text-muted-foreground">{key}</span>
								<span>{value}</span>
							</div>
						{/each}
					</TabsContent>
				</Tabs>
			</div>

			<div class="rounded-2xl border border-dashed border-border/80 bg-white p-4">
				<div class="flex items-center gap-2">
					<ShieldIcon class="size-4 text-primary" />
					<p class="text-sm font-medium text-foreground">Safety-aware routing</p>
				</div>
				<p class="mt-2 text-xs leading-5 text-muted-foreground">
					The shell communicates the guest limit model without pretending that custom execution is live.
				</p>
			</div>
		</CardContent>
	</Card>
</div>

<div class="mt-4 grid gap-4 lg:grid-cols-[1.2fr_0.8fr]">
	<Card class="border-border/80 shadow-sm">
		<CardHeader class="flex-row items-center justify-between gap-3">
			<div>
				<CardTitle>Workspace rails</CardTitle>
				<CardDescription>Collections, examples, and recent activity stay visible in the shell.</CardDescription>
			</div>
			<Badge variant="outline">Persistent</Badge>
		</CardHeader>
		<CardContent class="grid gap-3 sm:grid-cols-3">
			<div class="rounded-2xl border border-border/70 bg-panel-soft p-4">
				<div class="flex items-center gap-2">
					<InboxIcon class="size-4 text-primary" />
					<p class="text-sm font-medium text-foreground">Examples</p>
				</div>
				<p class="mt-2 text-xs leading-5 text-muted-foreground">JSONPlaceholder, GitHub, and weather demos.</p>
			</div>
			<div class="rounded-2xl border border-border/70 bg-panel-soft p-4">
				<div class="flex items-center gap-2">
					<LayoutGridIcon class="size-4 text-primary" />
					<p class="text-sm font-medium text-foreground">Collections</p>
				</div>
				<p class="mt-2 text-xs leading-5 text-muted-foreground">Saved groups will appear here after auth lands.</p>
			</div>
			<div class="rounded-2xl border border-border/70 bg-panel-soft p-4">
				<div class="flex items-center gap-2">
					<Clock3Icon class="size-4 text-primary" />
					<p class="text-sm font-medium text-foreground">History</p>
				</div>
				<p class="mt-2 text-xs leading-5 text-muted-foreground">Recent runs and response snapshots belong here.</p>
			</div>
		</CardContent>
	</Card>

	<Card class="border-border/80 bg-gradient-to-br from-emerald-50 to-white shadow-sm">
		<CardHeader class="gap-3">
			<div class="flex items-center justify-between gap-3">
				<div>
					<CardTitle>Guest lock state</CardTitle>
					<CardDescription>Visible module, restricted action.</CardDescription>
				</div>
				<Badge variant="secondary">Preview</Badge>
			</div>
		</CardHeader>
		<CardContent class="space-y-3">
			<div class="flex items-start gap-3 rounded-2xl border border-border/70 bg-white p-4">
				<LockIcon class="mt-0.5 size-4 text-muted-foreground" />
				<div>
					<p class="text-sm font-medium text-foreground">Custom target execution is locked</p>
					<p class="mt-1 text-xs leading-5 text-muted-foreground">
						Guests can inspect the full workspace, but only allowlisted endpoints should be runnable.
					</p>
				</div>
			</div>
			<Button variant="outline" class="w-full justify-between">
				<span>Open sign-in flow</span>
				<ArrowRightIcon class="size-4" />
			</Button>
		</CardContent>
	</Card>
</div>

<Separator class="my-4" />

<div class="grid gap-4 lg:grid-cols-[1.1fr_0.9fr]">
	<Card class="border-border/80 shadow-sm">
		<CardHeader class="gap-2">
			<CardTitle>Request metrics</CardTitle>
			<CardDescription>Lightweight shell-level summaries for the workspace.</CardDescription>
		</CardHeader>
		<CardContent class="grid gap-3 sm:grid-cols-3">
			<div class="rounded-2xl border border-border/70 bg-panel-soft p-4">
				<p class="text-xs uppercase tracking-[0.24em] text-muted-foreground">Today</p>
				<p class="mt-2 text-2xl font-semibold tracking-tight text-foreground">12</p>
				<p class="text-xs text-muted-foreground">guest-safe sends</p>
			</div>
			<div class="rounded-2xl border border-border/70 bg-panel-soft p-4">
				<p class="text-xs uppercase tracking-[0.24em] text-muted-foreground">Success</p>
				<p class="mt-2 text-2xl font-semibold tracking-tight text-foreground">98%</p>
				<p class="text-xs text-muted-foreground">preview responses</p>
			</div>
			<div class="rounded-2xl border border-border/70 bg-panel-soft p-4">
				<p class="text-xs uppercase tracking-[0.24em] text-muted-foreground">Quota</p>
				<p class="mt-2 text-2xl font-semibold tracking-tight text-foreground">8 left</p>
				<p class="text-xs text-muted-foreground">before cooldown</p>
			</div>
		</CardContent>
	</Card>

	<Card class="border-border/80 shadow-sm">
		<CardHeader class="gap-2">
			<CardTitle>Navigation intent</CardTitle>
			<CardDescription>The route structure matches the docs-led product map.</CardDescription>
		</CardHeader>
		<CardContent class="space-y-2">
			<div class="flex items-center justify-between rounded-2xl border border-border/70 bg-panel-soft px-3 py-2 text-sm">
				<span>/app</span>
				<span class="text-muted-foreground">Shell + workspace</span>
			</div>
			<div class="flex items-center justify-between rounded-2xl border border-border/70 bg-panel-soft px-3 py-2 text-sm">
				<span>/templates</span>
				<span class="text-muted-foreground">Example collections</span>
			</div>
			<div class="flex items-center justify-between rounded-2xl border border-border/70 bg-panel-soft px-3 py-2 text-sm">
				<span>/docs</span>
				<span class="text-muted-foreground">Quick start</span>
			</div>
		</CardContent>
	</Card>
</div>
