export default function HomePage() {
    return (
        <div className="flex h-screen">
            <div className="flex-1 flex-col">
                <div className="p-4">
                    <Chart />
                </div>
                <div className="p-4">
                    <Holdings />
                </div>
            </div>
            <div className="w-350 p-4 border-l-4 border-gray-200">
                <OrderBook />
            </div>
        </div>
    )
}