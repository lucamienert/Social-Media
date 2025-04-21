import Link from "next/link"
import IndexLayout from "./layouts/index-layout"
import { Button } from "@/components/ui/button"
import { ThemeToggle } from "@/components/theme-toggle"

export default function LandingPage() {
  return (
    <IndexLayout>
      <main className="min-h-screen flex flex-col">
        <nav className="w-full flex justify-between items-center p-6 border-b">
          <div className="text-2xl font-bold">Social Media</div>
          <div className="flex gap-4">
            <Link href="/auth/login">
              <Button variant="outline">Login</Button>
            </Link>
            <Link href="/auth/register">
              <Button>Register</Button>
            </Link>
          </div>
        </nav>

        <section className="flex-1 flex flex-col justify-center items-center text-center px-4">
          <h1 className="text-4xl font-bold mb-4">Connect. Share. Grow.</h1>
          <p className="text-lg text-muted-foreground max-w-md mb-6">
            Join Social Media
          </p>
          <Link href="/auth/register">
            <Button size="lg">Get Started</Button>
          </Link>
        </section>
        <ThemeToggle />
      </main>
    </IndexLayout>
  )
}
