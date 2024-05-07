import { OutlinedButton, OutlinedTextField } from "./material";

export default function Menu({}: {}) {
  return (
    <div className="flex flex-col gap-6 max-w-[230px] max-h-[380px]">
      <OutlinedTextField label="Title"></OutlinedTextField>
      <OutlinedTextField
        label="Description"
        type="textarea"
        rows={3}
        maxLength={100}
        className="max-h-[250px] min-h-[100px]"
      ></OutlinedTextField>
      <div>
        <label htmlFor="date" className="font-[Roboto]">
          Release date
        </label>
        <input id="date" type="date" />
      </div>
      <OutlinedButton>submit</OutlinedButton>
    </div>
  );
}
