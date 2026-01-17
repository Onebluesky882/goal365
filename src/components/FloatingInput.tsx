export type FloatingInputProps = {
  label: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  type?: string;
  id: string;
};

export const FloatingInput = ({
  label,
  value,
  onChange,
  type = "text",
  id,
}: FloatingInputProps) => {
  return (
    <div className="relative mb-4 h-14 bg-[#e8e8e8] rounded-t overflow-hidden">
      <input
        id={id}
        type={type}
        value={value}
        onChange={onChange}
        placeholder=" " // ห้ามลบ
        required
        className="peer w-full h-full px-4 pt-5 bg-transparent text-black focus:outline-none placeholder-transparent"
      />
      <label
        htmlFor={id}
        className="absolute left-4 top-1/2 -translate-y-1/2 scale-100 origin-left text-gray-500 transition-all duration-200 pointer-events-none
          peer-focus:-translate-y-4 peer-focus:scale-75
          peer-[:not(:placeholder-shown)]:-translate-y-4 peer-[:not(:placeholder-shown)]:scale-75"
      >
        {label}
      </label>
      <span className="absolute bottom-0 left-0 w-full h-px bg-black/20" />
      <span className="absolute bottom-0 left-0 w-full h-px scale-x-0 bg-black transition-transform duration-300 peer-focus:scale-x-100" />
    </div>
  );
};
