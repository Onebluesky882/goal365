"use client";
import CreatePlayerForm from "@/components/Player/CreatePlayer";
import { useAuth } from "@/GlobalContext/auth-provider";
import React, { useState } from "react";

export default function CreatePlayerContainer() {
  const { session } = useAuth();
  const [name, setName] = useState("");
  const [bio, setBio] = useState("");

  return (
    <CreatePlayerForm
      name={name}
      bio={bio}
      userId={session?.user?.id}
      onChangeName={setName}
      onChangeBio={setBio}
      onSuccess={() => {
        setName("");
        setBio("");
      }}
    />
  );
}
