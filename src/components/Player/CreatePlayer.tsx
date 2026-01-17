"use client";

import { playersApi } from "@/api/api";
import { handleError } from "@/lib/handleErrors";
import { useRouter } from "next/navigation";
import { toast } from "sonner";

type Props = {
  name: string;
  bio?: string;
  userId?: string;
  onChangeName: (v: string) => void;
  onChangeBio: (v: string) => void;
  onSuccess?: () => void;
};

export default function CreatePlayerForm({
  name,
  bio = "",
  userId,
  onChangeName,
  onChangeBio,
  onSuccess,
}: Props) {
  const router = useRouter();
  const handleCreate = async () => {
    if (!userId) {
      toast.error("กรุณา login ก่อน");
      return;
    }

    if (!name.trim()) {
      toast.error("กรุณากรอกชื่อ Player");
      return;
    }
    try {
      await playersApi.CreatePlayer({
        name,
        bio,
        userId,
      });

      toast.success("สร้าง Player สำเร็จ");
      onSuccess?.();
      router.push("/players");
    } catch (err) {
      handleError(err, "create player failed");
    }
  };

  return (
    <div className="rounded-md border p-2 border-gray-100 shadow-sm  m-2 mt-4">
      <h2 className="text-xl font-bold mb-4 text-gray-800">
        สร้าง Player ใหม่
      </h2>

      <input
        type="text"
        placeholder="ตั้งชื่อ"
        value={name}
        onChange={(e) => onChangeName(e.target.value)}
        className="w-full rounded-lg border px-4 py-2 mb-3
                   focus:outline-none focus:ring-2 focus:ring-blue-500"
      />

      <textarea
        placeholder="สเต็ปนอก ตอกเจ็บเจ็บ"
        value={bio}
        onChange={(e) => onChangeBio(e.target.value)}
        rows={3}
        className="w-full rounded-lg border px-4 py-2 mb-4 resize-none
                   focus:outline-none focus:ring-2 focus:ring-blue-500"
      />

      <button
        onClick={handleCreate}
        className="w-full rounded-lg bg-blue-600 text-white py-2 font-semibold
                   hover:bg-blue-700 transition"
      >
        Create Player
      </button>
    </div>
  );
}
