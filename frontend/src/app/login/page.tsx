import { LoginForm } from "@/components/login-form"
import IndexLayout from "../layouts/index-layout"
import { ThemeToggle } from "@/components/theme-toggle"

export default function Page() {
  return (
    <IndexLayout>
      <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
        <div className="w-full max-w-sm">
          <LoginForm />
          <ThemeToggle />
        </div>
      </div>
    </IndexLayout>
  )
}
