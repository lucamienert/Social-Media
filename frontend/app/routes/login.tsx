import type { Route } from "../+types/root";
import LoginForm from "~/Auth/LoginForm";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "New React Router App" },
    { name: "description", content: "Welcome to React Router!" },
  ];
}

export default function Login() {
  return <LoginForm />;
}