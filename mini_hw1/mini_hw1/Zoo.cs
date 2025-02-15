namespace mini_hw1;

public class Zoo
{
    private readonly VeterinaryClinic _clinic;
    private readonly List<IAnimal> _animals = new();
    private readonly List<IThing> _things = new();

    public Zoo(VeterinaryClinic clinic)
    {
        _clinic = clinic;
    }

    public bool AddAnimal(IAnimal animal)
    {
        if (_clinic.CheckHealth(animal))
        {
            _animals.Add(animal);
            Console.WriteLine($"{animal.Name} принято в зоопарк.");
            return true;
        }
        else
        {
            Console.WriteLine($"{animal.Name} не принято в зоопарк.");
            return false;
        }
    }

    public bool AddThing(IThing thing)
    {
        _things.Add(thing);
        Console.WriteLine($"Предмет {thing.Name} добавлен в инвентарь.");
        return true;
    }

    public void PrintReport()
    {
        Console.WriteLine("В данный момент в зоопарке");
        Console.WriteLine($"Животных: {_animals.Count}");
        Console.WriteLine($"Потребляется еды в день: {_animals.Sum(a => a.Food)} кг");
    }

    public void PrintContactZooAnimals()
    {
        Console.WriteLine("В контактном зоопарке могут быть следующие животные: ");
        foreach (var animal in _animals)
        {
            if (animal is IHerbo && animal.CanBeInContactZoo())
            {
                Console.WriteLine($"{animal.Name}, № {animal.Number}");
            }
        }
    }

    public void PrintInventory()
    {
        Console.WriteLine("Инвентарь:");
        foreach (var thing in _things)
        {
            Console.WriteLine($"{thing.Name}, № {thing.Number}");
        }
        foreach (var animal in _animals)
        {
            Console.WriteLine($"{animal.Name}, № {animal.Number}");
        }
    }
}