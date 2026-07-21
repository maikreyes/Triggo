export default function SelectedInput({label, id, options, value, onChange}:
    {label:string,id:string,options:string[], value:string, onChange: (value:string) => void }) {
    return (
        <div className="flex flex-col items-start justify-center gap-4">
            <label htmlFor={id} className="text-cyan-950">{label}</label>
            <select 
                id={id} 
                value={value}
                onChange={(e) => onChange(e.target.value)}
                className="bg-cyan-50 text-black">
                    <option value="" className="p-1">Seleccione un repositorio...</option>
                    {options.map((option) =>(
                        <option 
                            key={option} 
                            value={option} 
                            className="p-1"
                            >
                                {option}
                            </option>
                ))}
            </select>
        </div>
    )
}