import React from "react";

type User = {
  id: number;
  name: string;
  email: string;
};

type UsersProps = {
  users: User[];
};

export async function getServerSideProps() {
  const res = await fetch("http://backend:8080/users");
  const users = await res.json();

  return {
    props: { users },
  };
}

const Users: React.FC<UsersProps> = ({ users }) => {
  return (
    <div>
      <h1>ユーザー一覧</h1>
      <ul>
        {users.map((user) => (
          <li key={user.id}>
            {user.name} ({user.email})
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Users;
