import * as React from "react"
import { Outlet, createRootRoute } from "@tanstack/react-router"

import { ThemeProvider } from "@/components/theme-provider.tsx"

export const Route = createRootRoute({
  component: RootComponent,
})

function RootComponent() {
  return (
    <React.Fragment>
      <ThemeProvider>
        <div>Hello "__root"!</div>
        <Outlet />
      </ThemeProvider>
    </React.Fragment>
  )
}
