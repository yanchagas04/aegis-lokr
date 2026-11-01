export function Home() {
    return (
        <>
            <main className='flex flex-col min-h-screen w-full bg-linear-to-tr from-[#1e1e1e] to-[#06162a] justify-center items-center'>
                <img
                    src='/aegis-lokr.svg'
                    alt='Ícone Aegis Lokr'
                    className='mb-4 h-24 w-24 object-contain'
                />
                <h1 className='text-3xl text-white'>Aegis Lokr</h1>
            </main>
        </>
    )
}