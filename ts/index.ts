import { User } from "./types/user";

const greet = (user: { name: string }) => {
  console.log(`Hello, ${user.name}!`);
};

const user: User = {
  name: "Dan",
  age: 22,
};

greet(user);
