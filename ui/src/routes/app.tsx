import { Button } from "@/components/ui/button"
import { createFileRoute } from "@tanstack/react-router"
import { useState } from "react"

export const Route = createFileRoute("/app")({
  component: RouteComponent,
})

function RouteComponent() {
  const [res, setRes] = useState("")

  async function handlePing() {
    setRes("")
    const response = await (await fetch("/api/v1/ping")).json().catch()
    setRes(JSON.stringify(response))
  }
  return (
    <div className="flex flex-col gap-4">
      Hello "/app"!
      <Button onClick={() => handlePing()}>Click to call ping api</Button>
      <code>Result: {res}</code>
    </div>
  )
}
