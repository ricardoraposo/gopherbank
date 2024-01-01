import AmountBtn from "./AmountBtn";
import AmountInput from "./AmountInput";

function AmountDisplay() {

  return (
    <div className="flex flex-col items-center">
      <div className="flex gap-1 justify-center">
        <AmountInput />
      </div>
      <div className="flex gap-4 mt-8">
        <AmountBtn value={"5"} />
        <AmountBtn value={"10"} />
        <AmountBtn value={"20"} />
        <AmountBtn value={"50"} />
        <AmountBtn value={"100"} />
      </div>
    </div>
  )
}

export default AmountDisplay;
