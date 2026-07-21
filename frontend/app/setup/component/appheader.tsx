import Image from "next/image"

export default function AppHeader({imgSrc}:{imgSrc:string}) {
    return (
        <div className="flex flex-col items-center">
            <Image
            className="rounded-full m-4"
            src={imgSrc}
            alt="Logo"
            width={120}
            height={120}
            loading="eager"
            />
            <h1 className="text-3xl text-cyan-950">Triggo</h1>
        </div>
    )
}