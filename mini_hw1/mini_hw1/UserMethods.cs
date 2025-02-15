namespace mini_hw1;

public class UserMethods
{
    public static int ReadNumber(int k)
    {
        int n;
        Console.WriteLine($"Введите число от 1 до {k}");
        while (!int.TryParse(Console.ReadLine(), out n) || n > k || n < 1)
        {
            Console.WriteLine($"Вы ввели неверные данные!!! Введите число от 1 до {k}");
        }

        return n;
    }

    public static void AddAnimal(Zoo zoo, AnimalFactory animalFactory)
    {
        Console.WriteLine("Выберите животное:\n" +
                          "1. Тигр.\n" +
                          "2. Волк.\n" +
                          "3. Кролик.\n" +
                          "4. Обезьяна.");
        int n = ReadNumber(4);
        Console.WriteLine("Введиет имя: ");
        string name = Console.ReadLine();
        
        Console.WriteLine("Введите возраст. Животных старше 100 лет зоопарк не принимает: ");
        int age = ReadNumber(100);

        Console.WriteLine("Введите количество еды в день в кг. Учтите, что зоопарк не может обеспечить более 10кг еды в день для животного: ");
        int food = ReadNumber(10);
        Console.WriteLine("Введите уровень доброты/опасности: ");
        int index = ReadNumber(10);
        IAnimal animal = n switch
        {
            1 => animalFactory.CreateTiger(name, food, age, index),
            2 => animalFactory.CreateWolf(name, food, age, index),
            3 => animalFactory.CreateRabbit(name, food, age, index),
            4 => animalFactory.CreateMonkey(name, food, age, index),
            _ => null
        };
        if (animal != null)
        {
            zoo.AddAnimal(animal);
        }
    }
    public static void AddThing(Zoo zoo, ThingFactory thingFactory)
    {
        Console.WriteLine("Выберите предмет:\n" +
                          "1. Стол.\n" +
                          "2. Компьютер.\n");
        int n = ReadNumber(2);
        Console.WriteLine("Введиет название: ");
        string name = Console.ReadLine();
        IThing thing = n switch
        {
            1 => thingFactory.CreateTable(name),
            2 => thingFactory.CreateComputer(name),
            _ => null
        };
        if (thing != null)
        {
            zoo.AddThing(thing);
        }
    }
}