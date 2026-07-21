export default function LabeledInput({label, type, id, value, onChange}:
    {label:string, type:string, id:string, value:string, onChange(discordUrl:string):void }) {

    return (
        <div className="flex flex-col items-start justify-center gap-4">
            <label htmlFor={id} className="text-cyan-950">{label}</label>
            <input 
                type={type} 
                id={id} 
                value={value} 
                placeholder="Enter Discord Webhook Url..." 
                onChange={(e) => onChange(e.target.value)}
                className="bg-blue-50 rounded-md text-black w-full pl-1" 
            />
        </div>
    )

}