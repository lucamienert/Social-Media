import IndexLayout from "../layouts/index-layout"
import { ThemeToggle } from "@/components/theme-toggle"
import { RegisterForm } from "@/components/register-form"

export default function Page() {
  return (
    <IndexLayout>
      <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
        <div className="w-full max-w-sm">
          <RegisterForm />
          <ThemeToggle />
        </div>
      </div>
    </IndexLayout>
  )
}
