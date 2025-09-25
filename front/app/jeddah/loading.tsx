import { UixGlobalIconvcRfresh } from "../global/ui/server/uixGlobalIconvc";

export default function Page() {
  return (
    <div className="w-screen h-screen fixed top-0 flexctr bg-gradient-to-br from-sky-200 via-white to-emerald-100">
      <div className="text-3xl font-semibold text-slate-600 flexctr">
        <div>Loading</div>
        <div className="animate-spin">
          <UixGlobalIconvcRfresh bold={3} color="#45556c" size={2} />
        </div>
      </div>
    </div>
  );
}
