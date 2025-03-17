import {
  isRouteErrorResponse,
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
} from "react-router";

import type { Route } from "./+types/root";
import 'bootstrap/dist/css/bootstrap.min.css';
import "./app.css";
import Sidebar from "./components/Sidebar";
import type { ReactNode } from "react";
import { LoadingProvider } from "./Contexts/LoadingContext";
import Spinner from "./components/Spinner";
import { AuthProvider, useAuth } from "./Contexts/AuthContext";

// Link function for adding external resources like fonts
export const links: Route.LinksFunction = () => [
  { rel: "preconnect", href: "https://fonts.googleapis.com" },
  {
    rel: "preconnect",
    href: "https://fonts.gstatic.com",
    crossOrigin: "anonymous",
  },
  {
    rel: "stylesheet",
    href: "https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&display=swap",
  },
];

// Main layout that will wrap everything with the AuthProvider
export function Layout({ children }: { children: ReactNode }) {
  return (
    <AuthProvider>
      <HtmlLayout>{children}</HtmlLayout>
    </AuthProvider>
  );
}

// Layout containing the Sidebar and main content (depending on authentication)
function HtmlLayout({ children }: { children: ReactNode }) {
  const isAuthenticated = false// useAuth(); // Using the auth context

  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <meta name="description" content="A React application with sidebar navigation" />
        <Links />
      </head>
      <body>
        <LoadingProvider>
          {/* Only show the spinner if loading state is active */}
          <Spinner />
          <div className="d-flex">
            {isAuthenticated && <Sidebar />}
            {/* Main content */}
            <div className="flex-grow-1 p-3">
              {children}
            </div>
          </div>
          <ScrollRestoration />
          <Scripts />
        </LoadingProvider>
      </body>
    </html>
  );
}

// The App component that renders child routes
export default function App() {
  return <Outlet />;
}

// Error Boundary to handle route errors
export function ErrorBoundary({ error }: Route.ErrorBoundaryProps) {
  let message = "Oops!";
  let details = "An unexpected error occurred.";
  let stack: string | undefined;

  if (isRouteErrorResponse(error)) {
    message = error.status === 404 ? "404" : "Error";
    details =
      error.status === 404
        ? "The requested page could not be found."
        : error.statusText || details;
  } else if (import.meta.env.DEV && error instanceof Error) {
    details = error.message;
    stack = error.stack;
  }

  return (
    <main className="pt-16 p-4 container mx-auto">
      <h1>{message}</h1>
      <p>{details}</p>
      {stack && (
        <pre className="w-full p-4 overflow-x-auto">
          <code>{stack}</code>
        </pre>
      )}
    </main>
  );
}
