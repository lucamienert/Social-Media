import type { Route } from "../+types/root";
import RegisterForm from "~/Auth/RegisterForm";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "New React Router App" },
    { name: "description", content: "Welcome to React Router!" },
  ];
}

export default function Register() {
  return <RegisterForm />;
}