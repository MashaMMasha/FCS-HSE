using mini_hw1;
using System;
using System.Linq;
using System.Collections.Generic;
using Microsoft.Extensions.DependencyInjection;
internal class Program
{
    public static void Main(string[] args)
    {
        var serviceProvider = new ServiceCollection()
            .AddSingleton<VeterinaryClinic>()
            .AddSingleton<Zoo>()
            .AddSingleton<InventoryManager>()
            .AddSingleton<AnimalFactory>()
            .AddSingleton<ThingFactory>()
            .BuildServiceProvider();

        var veterinaryClinic = serviceProvider.GetService<VeterinaryClinic>();
        var zoo = serviceProvider.GetRequiredService<Zoo>();
        var inventoryManager = serviceProvider.GetRequiredService<InventoryManager>();
        var animalFactory = serviceProvider.GetRequiredService<AnimalFactory>();
        var thingFactory = serviceProvider.GetRequiredService<ThingFactory>();
        
        do
        {
            try
            {
                Console.WriteLine("Выберите число от 1 до 5:\n" +
                                  "1. Добавить животное.\n" +
                                  "2. Добавить предмет.\n" +
                                  "3. Вывести отчет по зоопарку.\n" +
                                  "4. Вывести животных для контактного зоопарка.\n" +
                                  "5. Вывести список инвентаря зоопарка.");
                int n = UserMethods.ReadNumber(5);
                switch (n)
                {
                    case 1:
                        UserMethods.AddAnimal(zoo, animalFactory);
                        break;
                    case 2:
                        UserMethods.AddThing(zoo, thingFactory);
                        break;
                    case 3:
                        zoo.PrintReport();
                        break;
                    case 4:
                        zoo.PrintContactZooAnimals();
                        break;
                    case 5:
                        zoo.PrintInventory();
                        break;
                    default:
                        Console.ForegroundColor = ConsoleColor.Red;
                        Console.WriteLine("Такого варианта нет!");
                        break;
                }
            }
            catch (Exception e)
            {
                Console.BackgroundColor = ConsoleColor.Red;
                Console.BackgroundColor = ConsoleColor.White;
                Console.WriteLine("Возникла ошибка при работе программы. Попробуйте еще раз :(" + e.Message);
            }

            Console.WriteLine("Нажмите enter чтобы продолжить, esc чтобы завершить работу программы.");
        } while (Console.ReadKey().Key != ConsoleKey.Escape);
    }
}