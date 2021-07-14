// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.Configstation
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("configstation");
        public static ImmutableArray<string> FavoritePlants { get; set; } = __config.GetObject<ImmutableArray<string>>("favoritePlants");

        /// <summary>
        /// omg my favorite sandwich
        /// </summary>
        public static Pulumi.Configstation.Config.Types.Sandwich? FavoriteSandwich { get; set; } = __config.GetObject<Pulumi.Configstation.Config.Types.Sandwich>("favoriteSandwich");

        public static bool? IsMember { get; set; } = __config.GetBoolean("isMember") ?? true;

        public static Types.Child? Kids { get; set; } = __config.GetObject<Types.Child>("kids");

        public static string? Name { get; set; } = __config.Get("name");

        public static int? NumberOfSheep { get; set; } = __config.GetInt32("numberOfSheep");

        /// <summary>
        /// This is a huge secret
        /// </summary>
        public static string? SecretCode { get; set; } = __config.Get("secretCode") ?? Utilities.GetEnv("SECRET_CODE", "MY_SUPER_SECRET_CODE");

        public static class Types
        {

             public class Sandwich
             {
                public string? Bread { get; set; } = null!;
                public ImmutableArray<string> Veggies { get; set; }
            }
        }
    }
}
