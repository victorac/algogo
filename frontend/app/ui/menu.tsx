export default function Menu({}: {}) {
  return (
    <ul className="flex flex-col gap-2">
      <li>
        <div className="relative">
          <input type="text" className="px-2 peer pt-6 pb-2"></input>
          <label className="absolute text-sm font-semibold left-2 top-1/2 -translate-y-1/2 z-10 px-1  peer-focus:top-0">
            Title
          </label>
        </div>
      </li>
    </ul>
  );
}
