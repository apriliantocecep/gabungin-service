import { Family } from "../types/family";

interface Props {
    family: Family | null;
}


const FamilyStructure = ({ family }: Props) => {

    if (!family) {
        return null
    }

    return (
        <div className="text-center">
            <div className={`inline-flex flex-col p-3 text-center ${family.gender == 'Male' ? 'bg-purple-800' : 'bg-pink-800'}`}>
                <p>{family.name}</p>
                <p>{family.order}</p>
            </div>
            <div className="py-4 flex justify-center gap-0 text-center">
                    {family.data && family.data.map((family, i) => {
                        return (
                            <FamilyStructure key={i} family={family} />
                        )
                    })}
                </div>
        </div>
    )
}

export default FamilyStructure