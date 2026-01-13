import React from "react";
type PageProps = {
  params: {
    playerNo: string;
  };
};

export default function Page({ params }: PageProps) {
  const { playerNo } = params; // Access the dynamic 'id'

  return <h1>player No ID: {playerNo}</h1>;
}
